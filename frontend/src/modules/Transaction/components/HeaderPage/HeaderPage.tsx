import styles from './HeaderPage.module.css';

type HeaderPageProps = {
  onClickUpload: () => void;
};

const HeaderPage = ({ onClickUpload }: HeaderPageProps) => {
  return (
    <div className={styles.header}>
      <div className={styles.title_wrapper}>
        <h1 className={styles.title}>Transaction Overview</h1>
        <p className={styles.description}>
          Upload your CSV file to process transactions, monitor your balance, and review any failed
          transactions that need attention.
        </p>
      </div>
      <button type="button" className={styles.button} onClick={onClickUpload}>
        Upload Transaction
      </button>
    </div>
  );
};

export default HeaderPage;
