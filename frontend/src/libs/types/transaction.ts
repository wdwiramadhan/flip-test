export type TransactionType = 'DEBIT' | 'CREDIT';

export type TransactionStatus = 'SUCCESS' | 'FAILED' | 'PENDING';

export interface Transaction {
  id: string;
  name: string;
  type: TransactionType;
  amount: number;
  status: TransactionStatus;
  description: string;
  transaction_date: string;
}
