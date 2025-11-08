import cn from 'classnames';
import { useEffect, useRef, useState } from 'react';
import { IconDocument, IconExclamationTriangle } from '@/components/icons';
import Modal from '@/components/Modal';
import styles from './UploadTransactionModal.module.css';

type UploadTransactionModalProps = {
  isOpen: boolean;
  isLoading?: boolean;
  onClose: () => void;
  onSubmit: (file: File) => void;
};

const UploadTransactionModal = ({
  isOpen,
  onClose,
  onSubmit,
  isLoading,
}: UploadTransactionModalProps) => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [isDragOver, setIsDragOver] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (!isOpen) {
      setSelectedFile(null);
      setError(null);
      setIsDragOver(false);
      if (fileInputRef.current) {
        fileInputRef.current.value = '';
      }
    }
  }, [isOpen]);

  const handleChooseFile = () => {
    fileInputRef.current?.click();
  };

  const validateFile = (file: File): boolean => {
    if (!file.name.endsWith('.csv')) {
      setError('Please upload a CSV file');
      return false;
    }
    if (file.size > 10 * 1024 * 1024) {
      setError('File size must be less than 10MB');
      return false;
    }
    setError(null);
    return true;
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file && validateFile(file)) {
      setSelectedFile(file);
    }
  };

  const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setIsDragOver(true);
  };

  const handleDragLeave = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setIsDragOver(false);
  };

  const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    setIsDragOver(false);

    const file = event.dataTransfer.files?.[0];
    if (file && validateFile(file)) {
      setSelectedFile(file);
    }
  };

  const handleSubmit = () => {
    if (selectedFile) {
      onSubmit(selectedFile);
    }
  };

  const handleClose = () => {
    setSelectedFile(null);
    setError(null);
    setIsDragOver(false);
    if (fileInputRef.current) {
      fileInputRef.current.value = '';
    }
    onClose();
  };

  const formatFileSize = (bytes: number): string => {
    if (bytes < 1024) return `${bytes} B`;
    if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(2)} KB`;
    return `${(bytes / (1024 * 1024)).toFixed(2)} MB`;
  };

  return (
    <Modal isOpen={isOpen} onClose={handleClose} title="Upload Transaction">
      <div className={styles.upload_container}>
        <input
          ref={fileInputRef}
          type="file"
          accept=".csv"
          onChange={handleFileChange}
          style={{ display: 'none' }}
        />
        {/* biome-ignore lint/a11y/noStaticElementInteractions: Drag and drop area is decorative, actual interaction handled by button inside */}
        <div
          className={cn(styles.drop_zone, {
            [styles.drop_zone_active]: isDragOver,
          })}
          onDragOver={handleDragOver}
          onDragLeave={handleDragLeave}
          onDrop={handleDrop}
          onClick={handleChooseFile}
          role="presentation"
        >
          <div className={styles.drop_zone_content}>
            <div className={styles.drop_zone_text}>
              <p className={styles.drop_zone_title}>
                {isDragOver ? 'Drop your CSV file here' : 'Drag and drop your CSV file here'}
              </p>
              <p className={styles.drop_zone_subtitle}>or click to browse</p>
            </div>
          </div>
        </div>

        {error && (
          <div className={styles.error_message}>
            <IconExclamationTriangle size={16} />
            <span>{error}</span>
          </div>
        )}

        {selectedFile && (
          <div className={styles.file_preview}>
            <div className={styles.file_icon}>
              <IconDocument size={24} />
            </div>
            <div className={styles.file_details}>
              <p className={styles.file_name}>{selectedFile.name}</p>
              <p className={styles.file_size}>{formatFileSize(selectedFile.size)}</p>
            </div>
            <button
              type="button"
              className={styles.remove_button}
              onClick={(e) => {
                e.stopPropagation();
                setSelectedFile(null);
                setError(null);
                if (fileInputRef.current) {
                  fileInputRef.current.value = '';
                }
              }}
              disabled={isLoading}
              aria-label="Remove file"
            >
              ×
            </button>
          </div>
        )}

        <div className={styles.actions}>
          <button
            type="button"
            className={styles.cancel_button}
            onClick={handleClose}
            disabled={isLoading}
          >
            Cancel
          </button>
          <button
            type="button"
            className={styles.submit_button}
            onClick={handleSubmit}
            disabled={!selectedFile || isLoading}
          >
            {isLoading ? 'Uploading...' : 'Upload Transaction'}
          </button>
        </div>
      </div>
    </Modal>
  );
};

export default UploadTransactionModal;
