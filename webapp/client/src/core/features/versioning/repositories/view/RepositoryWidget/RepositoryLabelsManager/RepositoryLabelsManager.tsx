import React from 'react';

import { IRepository } from 'core/shared/models/Versioning/Repository';
import TagsManager from 'core/shared/view/domain/TagsManager/TagsManager';
import { hasAccessToAction } from 'models/EntitiesActions';

import styles from './RepositoryLabelsManager.module.css';
import {
  useAddLabelMutation,
  useDeleteLabelMutation,
} from '../../../store/labelsManager/labelsManager';

interface ILocalProps {
  repository: IRepository;
}

type AllProps = ILocalProps;

const RepositoryLabelsManager: React.FC<AllProps> = ({ repository }) => {
  const [addLabel, addingLabel] = useAddLabelMutation();
  const [deleteLabel, deletingLabel] = useDeleteLabelMutation();

  const isAvailableEditing = hasAccessToAction('update', repository);

  return (
    <div className={styles.root}>
      <TagsManager
        tags={repository.labels}
        isDraggableTags={false}
        isAvailableTagAdding={isAvailableEditing}
        isRemovableTags={isAvailableEditing}
        isShowPlaceholder={repository.labels.length === 0}
        isUpdating={addingLabel.isRequesting || deletingLabel.isRequesting}
        onAddTag={label => addLabel({ id: repository.id, label })}
        onRemoveTag={label => deleteLabel({ id: repository.id, label })}
        tagWordReplacer="Label"
      />
    </div>
  );
};

export default RepositoryLabelsManager;
