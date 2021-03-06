package ai.verta.blobs.dataset

import ai.verta.client._
import ai.verta.blobs._
import ai.verta.swagger.client.HttpException
import ai.verta.swagger._public.modeldb.versioning.model.VersioningSetTagRequestResponse

import scala.concurrent.ExecutionContext
import scala.language.reflectiveCalls
import scala.util.{Try, Success, Failure}

import org.scalatest.FunSuite
import org.scalatest.Assertions._

import com.amazonaws.regions.Regions;
import com.amazonaws.services.s3._
import com.amazonaws.services.s3.model._
import collection.JavaConverters._

import scala.collection.mutable.HashSet

class TestS3 extends FunSuite {
  def fixture =
    new {
      val testfilePath = "s3://verta-starter/census-test.csv"
      val testfileLoc = S3Location(testfilePath).get

      val testfilePath2 = "s3://verta-starter/census-train.csv"
      val testfileLoc2 = S3Location(testfilePath2).get

      val testdirPath = "s3://verta-versioned-bucket/data/"
      val testdirLoc = S3Location(testdirPath).get

      val bucketLoc = S3Location("s3://verta-starter").get
    }

  def versioned_fixture = new {
    // for versioning tests:
    val s3: AmazonS3 = AmazonS3ClientBuilder.standard().build()
    val filepath = "s3://verta-versioned-bucket/data/census-train.csv"

    val versionListing = s3.listVersions("verta-versioned-bucket", "data/census-train.csv")
    val oldVersionId = versionListing.getVersionSummaries().asScala.toList
      .filter((version: S3VersionSummary) => !version.getKey().endsWith("/")) // not a folder
      .filter(!_.isLatest()).head.getVersionId()
    val newVersionId = versionListing.getVersionSummaries().asScala.toList
      .filter((version: S3VersionSummary) => !version.getKey().endsWith("/")) // not a folder
      .filter(_.isLatest()).head.getVersionId()

    val oldVersonLoc = S3Location(filepath, Some(oldVersionId)).get
    val newVersionLoc = S3Location(filepath, Some(newVersionId)).get
  }

  test("S3Location should correctly determine bucket name and key") {
    val f = fixture
    assert(f.testfileLoc.bucketName.equals("verta-starter"))
    assert(f.testfileLoc.key.get.equals("census-test.csv"))

    assert(f.bucketLoc.bucketName.equals("verta-starter"))
    assert(f.bucketLoc.key.isEmpty)
  }

  test("S3Location construction should fail when encountered non-S3 path") {
    val s3LocAttempt = S3Location("http://verta-starter/census-train.csv")
    assert(s3LocAttempt.isFailure)
    assert(s3LocAttempt match {case Failure(e) => e.getMessage contains "Illegal path; must be an S3 location"})
  }

  test("S3 blob should retrieve the file (i.e with key) correctly") {
    val f = fixture
    val s3Blob = S3(f.testfileLoc).get

    val s3File = s3Blob.getMetadata(f.testfilePath).get
    TestMetadata.assertMetadata(s3Blob.getMetadata(f.testfilePath).get, f.testfilePath)
    assert(s3Blob.listPaths equals List(f.testfilePath))
  }

  test("S3 blob should retrieve a folder correctly") {
    val f = fixture
    val s3Blob = S3(List(f.testdirLoc)).get

    val filePath = "s3://verta-versioned-bucket/data/census-train.csv"

    TestMetadata.assertMetadata(s3Blob.getMetadata(filePath).get, filePath)
    assert(s3Blob.getMetadata(f.testfilePath).isEmpty)
    assert(s3Blob.listPaths contains filePath)
  }

  test("S3 blob should retrieve the entire bucket correctly") {
    val f = fixture
    val s3Blob = S3(List(f.bucketLoc)).get

    TestMetadata.assertMetadata(s3Blob.getMetadata(f.testfilePath).get, f.testfilePath)
    TestMetadata.assertMetadata(s3Blob.getMetadata(f.testfilePath2).get, f.testfilePath2)
  }

  // testdir/testfile has two versions
  test("S3 should retrieve the correct version") {
    val f = fixture
    val vf = versioned_fixture

    // default should be the latest version
    val s3BlobDefault = S3(S3Location(vf.filepath).get).get
    var s3BlobLatest = S3(vf.newVersionLoc).get

    assert(s3BlobDefault equals s3BlobLatest)

    // get older version
    val s3BlobOld = S3(vf.oldVersonLoc).get
    assert(!s3BlobDefault.equals(s3BlobOld))
  }

  test("S3 should retrieve multiple keys correctly") {
    val f = fixture
    val s3Blob = S3(List(f.testfileLoc, f.testfileLoc2)).get

    TestMetadata.assertMetadata(s3Blob.getMetadata(f.testfilePath).get, f.testfilePath)
    TestMetadata.assertMetadata(s3Blob.getMetadata(f.testfilePath2).get, f.testfilePath2)
  }

  test("S3 should not have duplicate paths") {
    val f = fixture
    val s3Blob1 = S3(List(f.testfileLoc, f.testfileLoc, f.testfileLoc2, f.bucketLoc)).get

    val s3Blob2 = S3(f.bucketLoc).get
    assert(s3Blob1 equals s3Blob2)
  }

  test("S3 blob construction should fail when an invalid path is passed") {
    val f = fixture
    val invalidPath = "s3://verta-starter/no-such-file"
    val s3LocInvalid = S3Location(invalidPath).get

    val s3Blob = S3(List(s3LocInvalid, f.testfileLoc))

    assert(s3Blob.isFailure)
    assert(s3Blob match {case Failure(e) => e.getMessage contains "Not Found"})
  }

  ignore("S3 should produce similar hash and file size to PathBlob") {
    // This test no longer applies, since local path dataset and S3 dataset does not match
    val f = fixture
    val s3Blob = S3(List(f.testfileLoc2)).get
    val s3Metadata = s3Blob.getMetadata(f.testfilePath2).get

    val workingDir = System.getProperty("user.dir")
    val testDir = workingDir + "/src/test/scala/ai/verta/blobs/testdir/testsubdir/testfile2"
    val pathBlob = PathBlob(List(testDir)).get
    val pathBlobMetadata = pathBlob.getMetadata(testDir).get

    assert(s3Metadata.md5 equals pathBlobMetadata.md5)
    assert(s3Metadata.size equals pathBlobMetadata.size)
  }

  test("Reducing S3 blobs should retain the contents of both") {
    val f = fixture
    val s3Blob = S3(List(f.testfileLoc, f.testfileLoc2)).get
    val s3BlobCombined = S3.reduce(S3(f.testfileLoc).get, S3(f.testfileLoc2).get).get

    assert(s3Blob equals s3BlobCombined)
    assert(s3BlobCombined.listPaths.toSet equals Set(f.testfilePath, f.testfilePath2))
  }

  test("Reducing S3 blobs with conflicting contents should fail") {
    val vf = versioned_fixture

    val s3BlobDefault = S3(vf.newVersionLoc).get

    // get older version
    val s3BlobOld = S3(vf.oldVersonLoc).get

    val combineAttempt = S3.reduce(s3BlobDefault, s3BlobOld)
    assert(combineAttempt.isFailure)
    assert(combineAttempt match {case Failure(e) => e.getMessage contains "conflicting entries"})
  }
}
