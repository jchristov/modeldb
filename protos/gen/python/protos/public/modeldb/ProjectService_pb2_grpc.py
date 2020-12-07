# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from ..modeldb import CommonService_pb2 as modeldb_dot_CommonService__pb2
from ..modeldb import ProjectService_pb2 as modeldb_dot_ProjectService__pb2


class ProjectServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.createProject = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/createProject',
        request_serializer=modeldb_dot_ProjectService__pb2.CreateProject.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.CreateProject.Response.FromString,
        )
    self.getProjects = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getProjects',
        request_serializer=modeldb_dot_ProjectService__pb2.GetProjects.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.GetProjects.Response.FromString,
        )
    self.getPublicProjects = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getPublicProjects',
        request_serializer=modeldb_dot_ProjectService__pb2.GetPublicProjects.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.GetPublicProjects.Response.FromString,
        )
    self.getProjectById = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getProjectById',
        request_serializer=modeldb_dot_ProjectService__pb2.GetProjectById.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.GetProjectById.Response.FromString,
        )
    self.getProjectByName = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getProjectByName',
        request_serializer=modeldb_dot_ProjectService__pb2.GetProjectByName.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.GetProjectByName.Response.FromString,
        )
    self.deleteProject = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/deleteProject',
        request_serializer=modeldb_dot_ProjectService__pb2.DeleteProject.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.DeleteProject.Response.FromString,
        )
    self.updateProjectName = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/updateProjectName',
        request_serializer=modeldb_dot_ProjectService__pb2.UpdateProjectName.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.UpdateProjectName.Response.FromString,
        )
    self.updateProjectDescription = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/updateProjectDescription',
        request_serializer=modeldb_dot_ProjectService__pb2.UpdateProjectDescription.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.UpdateProjectDescription.Response.FromString,
        )
    self.addProjectTags = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/addProjectTags',
        request_serializer=modeldb_dot_ProjectService__pb2.AddProjectTags.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.AddProjectTags.Response.FromString,
        )
    self.getProjectTags = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getProjectTags',
        request_serializer=modeldb_dot_CommonService__pb2.GetTags.SerializeToString,
        response_deserializer=modeldb_dot_CommonService__pb2.GetTags.Response.FromString,
        )
    self.deleteProjectTags = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/deleteProjectTags',
        request_serializer=modeldb_dot_ProjectService__pb2.DeleteProjectTags.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjectTags.Response.FromString,
        )
    self.addProjectTag = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/addProjectTag',
        request_serializer=modeldb_dot_ProjectService__pb2.AddProjectTag.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.AddProjectTag.Response.FromString,
        )
    self.deleteProjectTag = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/deleteProjectTag',
        request_serializer=modeldb_dot_ProjectService__pb2.DeleteProjectTag.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjectTag.Response.FromString,
        )
    self.updateProjectAttributes = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/updateProjectAttributes',
        request_serializer=modeldb_dot_ProjectService__pb2.UpdateProjectAttributes.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.UpdateProjectAttributes.Response.FromString,
        )
    self.getProjectAttributes = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getProjectAttributes',
        request_serializer=modeldb_dot_CommonService__pb2.GetAttributes.SerializeToString,
        response_deserializer=modeldb_dot_CommonService__pb2.GetAttributes.Response.FromString,
        )
    self.addProjectAttributes = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/addProjectAttributes',
        request_serializer=modeldb_dot_ProjectService__pb2.AddProjectAttributes.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.AddProjectAttributes.Response.FromString,
        )
    self.deleteProjectAttributes = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/deleteProjectAttributes',
        request_serializer=modeldb_dot_ProjectService__pb2.DeleteProjectAttributes.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjectAttributes.Response.FromString,
        )
    self.logProjectCodeVersion = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/logProjectCodeVersion',
        request_serializer=modeldb_dot_ProjectService__pb2.LogProjectCodeVersion.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.LogProjectCodeVersion.Response.FromString,
        )
    self.getProjectCodeVersion = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getProjectCodeVersion',
        request_serializer=modeldb_dot_ProjectService__pb2.GetProjectCodeVersion.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.GetProjectCodeVersion.Response.FromString,
        )
    self.verifyConnection = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/verifyConnection',
        request_serializer=modeldb_dot_ProjectService__pb2.Empty.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.VerifyConnectionResponse.FromString,
        )
    self.deepCopyProject = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/deepCopyProject',
        request_serializer=modeldb_dot_ProjectService__pb2.DeepCopyProject.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.DeepCopyProject.Response.FromString,
        )
    self.getSummary = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getSummary',
        request_serializer=modeldb_dot_ProjectService__pb2.GetSummary.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.GetSummary.Response.FromString,
        )
    self.setProjectReadme = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/setProjectReadme',
        request_serializer=modeldb_dot_ProjectService__pb2.SetProjectReadme.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.SetProjectReadme.Response.FromString,
        )
    self.getProjectReadme = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getProjectReadme',
        request_serializer=modeldb_dot_ProjectService__pb2.GetProjectReadme.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.GetProjectReadme.Response.FromString,
        )
    self.setProjectVisibility = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/setProjectVisibility',
        request_serializer=modeldb_dot_ProjectService__pb2.SetProjectVisibility.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.SetProjectVisibility.Response.FromString,
        )
    self.setProjectShortName = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/setProjectShortName',
        request_serializer=modeldb_dot_ProjectService__pb2.SetProjectShortName.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.SetProjectShortName.Response.FromString,
        )
    self.getProjectShortName = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getProjectShortName',
        request_serializer=modeldb_dot_ProjectService__pb2.GetProjectShortName.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.GetProjectShortName.Response.FromString,
        )
    self.getUrlForArtifact = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getUrlForArtifact',
        request_serializer=modeldb_dot_CommonService__pb2.GetUrlForArtifact.SerializeToString,
        response_deserializer=modeldb_dot_CommonService__pb2.GetUrlForArtifact.Response.FromString,
        )
    self.findProjects = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/findProjects',
        request_serializer=modeldb_dot_ProjectService__pb2.FindProjects.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.FindProjects.Response.FromString,
        )
    self.logArtifacts = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/logArtifacts',
        request_serializer=modeldb_dot_ProjectService__pb2.LogProjectArtifacts.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.LogProjectArtifacts.Response.FromString,
        )
    self.getArtifacts = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/getArtifacts',
        request_serializer=modeldb_dot_CommonService__pb2.GetArtifacts.SerializeToString,
        response_deserializer=modeldb_dot_CommonService__pb2.GetArtifacts.Response.FromString,
        )
    self.deleteArtifact = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/deleteArtifact',
        request_serializer=modeldb_dot_ProjectService__pb2.DeleteProjectArtifact.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjectArtifact.Response.FromString,
        )
    self.deleteProjects = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/deleteProjects',
        request_serializer=modeldb_dot_ProjectService__pb2.DeleteProjects.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjects.Response.FromString,
        )
    self.setProjectWorkspace = channel.unary_unary(
        '/ai.verta.modeldb.ProjectService/setProjectWorkspace',
        request_serializer=modeldb_dot_ProjectService__pb2.SetProjectWorkspace.SerializeToString,
        response_deserializer=modeldb_dot_ProjectService__pb2.SetProjectWorkspace.Response.FromString,
        )


class ProjectServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def createProject(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getProjects(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getPublicProjects(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getProjectById(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getProjectByName(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def deleteProject(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def updateProjectName(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def updateProjectDescription(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def addProjectTags(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getProjectTags(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def deleteProjectTags(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def addProjectTag(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def deleteProjectTag(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def updateProjectAttributes(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getProjectAttributes(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def addProjectAttributes(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def deleteProjectAttributes(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def logProjectCodeVersion(self, request, context):
    """code version
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getProjectCodeVersion(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def verifyConnection(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def deepCopyProject(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getSummary(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def setProjectReadme(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getProjectReadme(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def setProjectVisibility(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def setProjectShortName(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getProjectShortName(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getUrlForArtifact(self, request, context):
    """artifacts
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def findProjects(self, request, context):
    """queries
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def logArtifacts(self, request, context):
    """artifacts
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def getArtifacts(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def deleteArtifact(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def deleteProjects(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def setProjectWorkspace(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_ProjectServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'createProject': grpc.unary_unary_rpc_method_handler(
          servicer.createProject,
          request_deserializer=modeldb_dot_ProjectService__pb2.CreateProject.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.CreateProject.Response.SerializeToString,
      ),
      'getProjects': grpc.unary_unary_rpc_method_handler(
          servicer.getProjects,
          request_deserializer=modeldb_dot_ProjectService__pb2.GetProjects.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.GetProjects.Response.SerializeToString,
      ),
      'getPublicProjects': grpc.unary_unary_rpc_method_handler(
          servicer.getPublicProjects,
          request_deserializer=modeldb_dot_ProjectService__pb2.GetPublicProjects.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.GetPublicProjects.Response.SerializeToString,
      ),
      'getProjectById': grpc.unary_unary_rpc_method_handler(
          servicer.getProjectById,
          request_deserializer=modeldb_dot_ProjectService__pb2.GetProjectById.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.GetProjectById.Response.SerializeToString,
      ),
      'getProjectByName': grpc.unary_unary_rpc_method_handler(
          servicer.getProjectByName,
          request_deserializer=modeldb_dot_ProjectService__pb2.GetProjectByName.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.GetProjectByName.Response.SerializeToString,
      ),
      'deleteProject': grpc.unary_unary_rpc_method_handler(
          servicer.deleteProject,
          request_deserializer=modeldb_dot_ProjectService__pb2.DeleteProject.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.DeleteProject.Response.SerializeToString,
      ),
      'updateProjectName': grpc.unary_unary_rpc_method_handler(
          servicer.updateProjectName,
          request_deserializer=modeldb_dot_ProjectService__pb2.UpdateProjectName.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.UpdateProjectName.Response.SerializeToString,
      ),
      'updateProjectDescription': grpc.unary_unary_rpc_method_handler(
          servicer.updateProjectDescription,
          request_deserializer=modeldb_dot_ProjectService__pb2.UpdateProjectDescription.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.UpdateProjectDescription.Response.SerializeToString,
      ),
      'addProjectTags': grpc.unary_unary_rpc_method_handler(
          servicer.addProjectTags,
          request_deserializer=modeldb_dot_ProjectService__pb2.AddProjectTags.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.AddProjectTags.Response.SerializeToString,
      ),
      'getProjectTags': grpc.unary_unary_rpc_method_handler(
          servicer.getProjectTags,
          request_deserializer=modeldb_dot_CommonService__pb2.GetTags.FromString,
          response_serializer=modeldb_dot_CommonService__pb2.GetTags.Response.SerializeToString,
      ),
      'deleteProjectTags': grpc.unary_unary_rpc_method_handler(
          servicer.deleteProjectTags,
          request_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjectTags.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.DeleteProjectTags.Response.SerializeToString,
      ),
      'addProjectTag': grpc.unary_unary_rpc_method_handler(
          servicer.addProjectTag,
          request_deserializer=modeldb_dot_ProjectService__pb2.AddProjectTag.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.AddProjectTag.Response.SerializeToString,
      ),
      'deleteProjectTag': grpc.unary_unary_rpc_method_handler(
          servicer.deleteProjectTag,
          request_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjectTag.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.DeleteProjectTag.Response.SerializeToString,
      ),
      'updateProjectAttributes': grpc.unary_unary_rpc_method_handler(
          servicer.updateProjectAttributes,
          request_deserializer=modeldb_dot_ProjectService__pb2.UpdateProjectAttributes.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.UpdateProjectAttributes.Response.SerializeToString,
      ),
      'getProjectAttributes': grpc.unary_unary_rpc_method_handler(
          servicer.getProjectAttributes,
          request_deserializer=modeldb_dot_CommonService__pb2.GetAttributes.FromString,
          response_serializer=modeldb_dot_CommonService__pb2.GetAttributes.Response.SerializeToString,
      ),
      'addProjectAttributes': grpc.unary_unary_rpc_method_handler(
          servicer.addProjectAttributes,
          request_deserializer=modeldb_dot_ProjectService__pb2.AddProjectAttributes.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.AddProjectAttributes.Response.SerializeToString,
      ),
      'deleteProjectAttributes': grpc.unary_unary_rpc_method_handler(
          servicer.deleteProjectAttributes,
          request_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjectAttributes.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.DeleteProjectAttributes.Response.SerializeToString,
      ),
      'logProjectCodeVersion': grpc.unary_unary_rpc_method_handler(
          servicer.logProjectCodeVersion,
          request_deserializer=modeldb_dot_ProjectService__pb2.LogProjectCodeVersion.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.LogProjectCodeVersion.Response.SerializeToString,
      ),
      'getProjectCodeVersion': grpc.unary_unary_rpc_method_handler(
          servicer.getProjectCodeVersion,
          request_deserializer=modeldb_dot_ProjectService__pb2.GetProjectCodeVersion.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.GetProjectCodeVersion.Response.SerializeToString,
      ),
      'verifyConnection': grpc.unary_unary_rpc_method_handler(
          servicer.verifyConnection,
          request_deserializer=modeldb_dot_ProjectService__pb2.Empty.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.VerifyConnectionResponse.SerializeToString,
      ),
      'deepCopyProject': grpc.unary_unary_rpc_method_handler(
          servicer.deepCopyProject,
          request_deserializer=modeldb_dot_ProjectService__pb2.DeepCopyProject.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.DeepCopyProject.Response.SerializeToString,
      ),
      'getSummary': grpc.unary_unary_rpc_method_handler(
          servicer.getSummary,
          request_deserializer=modeldb_dot_ProjectService__pb2.GetSummary.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.GetSummary.Response.SerializeToString,
      ),
      'setProjectReadme': grpc.unary_unary_rpc_method_handler(
          servicer.setProjectReadme,
          request_deserializer=modeldb_dot_ProjectService__pb2.SetProjectReadme.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.SetProjectReadme.Response.SerializeToString,
      ),
      'getProjectReadme': grpc.unary_unary_rpc_method_handler(
          servicer.getProjectReadme,
          request_deserializer=modeldb_dot_ProjectService__pb2.GetProjectReadme.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.GetProjectReadme.Response.SerializeToString,
      ),
      'setProjectVisibility': grpc.unary_unary_rpc_method_handler(
          servicer.setProjectVisibility,
          request_deserializer=modeldb_dot_ProjectService__pb2.SetProjectVisibility.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.SetProjectVisibility.Response.SerializeToString,
      ),
      'setProjectShortName': grpc.unary_unary_rpc_method_handler(
          servicer.setProjectShortName,
          request_deserializer=modeldb_dot_ProjectService__pb2.SetProjectShortName.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.SetProjectShortName.Response.SerializeToString,
      ),
      'getProjectShortName': grpc.unary_unary_rpc_method_handler(
          servicer.getProjectShortName,
          request_deserializer=modeldb_dot_ProjectService__pb2.GetProjectShortName.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.GetProjectShortName.Response.SerializeToString,
      ),
      'getUrlForArtifact': grpc.unary_unary_rpc_method_handler(
          servicer.getUrlForArtifact,
          request_deserializer=modeldb_dot_CommonService__pb2.GetUrlForArtifact.FromString,
          response_serializer=modeldb_dot_CommonService__pb2.GetUrlForArtifact.Response.SerializeToString,
      ),
      'findProjects': grpc.unary_unary_rpc_method_handler(
          servicer.findProjects,
          request_deserializer=modeldb_dot_ProjectService__pb2.FindProjects.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.FindProjects.Response.SerializeToString,
      ),
      'logArtifacts': grpc.unary_unary_rpc_method_handler(
          servicer.logArtifacts,
          request_deserializer=modeldb_dot_ProjectService__pb2.LogProjectArtifacts.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.LogProjectArtifacts.Response.SerializeToString,
      ),
      'getArtifacts': grpc.unary_unary_rpc_method_handler(
          servicer.getArtifacts,
          request_deserializer=modeldb_dot_CommonService__pb2.GetArtifacts.FromString,
          response_serializer=modeldb_dot_CommonService__pb2.GetArtifacts.Response.SerializeToString,
      ),
      'deleteArtifact': grpc.unary_unary_rpc_method_handler(
          servicer.deleteArtifact,
          request_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjectArtifact.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.DeleteProjectArtifact.Response.SerializeToString,
      ),
      'deleteProjects': grpc.unary_unary_rpc_method_handler(
          servicer.deleteProjects,
          request_deserializer=modeldb_dot_ProjectService__pb2.DeleteProjects.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.DeleteProjects.Response.SerializeToString,
      ),
      'setProjectWorkspace': grpc.unary_unary_rpc_method_handler(
          servicer.setProjectWorkspace,
          request_deserializer=modeldb_dot_ProjectService__pb2.SetProjectWorkspace.FromString,
          response_serializer=modeldb_dot_ProjectService__pb2.SetProjectWorkspace.Response.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'ai.verta.modeldb.ProjectService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
