import type { Transaction } from '@/libs/types/transaction';
import cn from 'classnames';
import { format } from 'date-fns';
import { IconExclamationTriangle, IconXMark } from '@/components/icons';
import { formatCurrency } from '../../utils/helpers';
import styles from './TransactionList.module.css';

type TransactionListProps = {
  transactions: Transaction[];
};

const TransactionList = ({ transactions }: TransactionListProps) => {
  if (transactions?.length === 0) {
    return (
      <div className={styles.empty_state}>
        <p>No transaction issues found</p>
      </div>
    );
  }

  return (
    <div className={styles.table_container}>
      <h2 className={styles.table_title}>Transaction Issues</h2>
      <div className={styles.table_wrapper}>
        <table className={styles.table}>
          <thead>
            <tr>
              <th>Date & Time</th>
              <th>Name</th>
              <th>Type</th>
              <th style={{ textAlign: 'right' }}>Amount</th>
              <th>Status</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {transactions.map((transaction) => (
              <tr key={transaction.id}>
                <td className={styles.date_cell}>
                  {format(new Date(transaction.transaction_date), 'dd MMM yyyy HH:mm:ss')}
                </td>
                <td className={styles.name_cell}>{transaction.name}</td>
                <td>
                  <div
                    className={cn(styles.type_badge, {
                      [styles.type_debit]: transaction.type === 'DEBIT',
                      [styles.type_credit]: transaction.type === 'CREDIT',
                    })}
                  >
                    {transaction.type}
                  </div>
                </td>
                <td className={styles.amount_cell}>{formatCurrency(transaction.amount)}</td>
                <td>
                  <div
                    className={cn(styles.status_badge, {
                      [styles.status_failed]: transaction.status === 'FAILED',
                      [styles.status_pending]: transaction.status === 'PENDING',
                    })}
                  >
                    {transaction.status === 'FAILED' ? (
                      <IconXMark size={16} />
                    ) : (
                      <IconExclamationTriangle size={16} />
                    )}
                    <span>{transaction.status}</span>
                  </div>
                </td>
                <td className={styles.description_cell}>{transaction.description}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default TransactionList;
