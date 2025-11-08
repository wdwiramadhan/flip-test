import cn from 'classnames';
import { formatCurrency } from '../../utils/helpers';
import styles from './Balance.module.css';

type BalanceProps = {
  balance: number;
};

const Balance = ({ balance }: BalanceProps) => {
  return (
    <div className={styles.card}>
      <p className={styles.label}>Current Balance</p>
      <div
        className={cn(styles.balance, {
          [styles.balance_negative]: balance < 0
        })}
      >
        {formatCurrency(balance)}
      </div>
    </div>
  );
};

export default Balance;
