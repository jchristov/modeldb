import * as React from 'react';

import DefaultMatchRemoteData from 'core/shared/view/elements/MatchRemoteDataComponents/DefaultMatchRemoteData';
import {
  IRepository,
  IBranchesAndTags,
} from 'core/shared/models/Versioning/Repository';

import { useRepositoryBranchesAndTagsQuery } from '../../../store/repositoryBranchesAndTags/useRepositoryBranchesAndTags';

type WrapperOwnProps = { repository: IRepository };

export default function withLoadingRequiredData<Props extends IBranchesAndTags>(
  WrappedComponent: React.ComponentType<Props>
): React.FC<Omit<Props, keyof IBranchesAndTags> & WrapperOwnProps> {
  function Wrapper(
    props: Omit<Props, keyof IBranchesAndTags> & WrapperOwnProps
  ) {
    const { communication, data } = useRepositoryBranchesAndTagsQuery({
      repositoryId: props.repository.id,
    });

    return (
      <DefaultMatchRemoteData communication={communication} data={data}>
        {loadedData => {
          return (
            <WrappedComponent
              {...props as any}
              branches={loadedData.branches}
              tags={loadedData.tags}
            />
          );
        }}
      </DefaultMatchRemoteData>
    );
  }
  return Wrapper;
}
