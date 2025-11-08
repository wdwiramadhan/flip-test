import { useQueryClient } from '@tanstack/react-query';
import { useState } from 'react';
import Container from '@/components/Container';
import { RQ_BALANCE, RQ_UNSUCCESSFUL_TRANSACTIONS } from '@/libs/constant';
import useBalanceQuery from '@/libs/hooks/useBalanceQuery';
import useUnsuccessfulTransactionQuery from '@/libs/hooks/useUnsuccessfulTransactionQuery';
import useUploadTransactionMutation from '@/libs/hooks/useUploadTransactionMutation';
import Balance from './components/Balance';
import HeaderPage from './components/HeaderPage';
import TransactionList from './components/TransactionList';
import UploadTransactionModal from './components/UploadTransactionModal';

const Transaction = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const queryClient = useQueryClient();
  const unsuccessfulTransactionQuery = useUnsuccessfulTransactionQuery();
  const balanceQuery = useBalanceQuery();
  const uploadTransactionMutation = useUploadTransactionMutation();

  const handleSubmitCsvTransaction = (file: File) => {
    const formData = new FormData();
    formData.append('file', file);
    uploadTransactionMutation.mutate(formData, {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: [RQ_BALANCE] });
        queryClient.invalidateQueries({ queryKey: [RQ_UNSUCCESSFUL_TRANSACTIONS] });
        setIsModalOpen(false);
      },
    });
  };

  return (
    <Container style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>
      <UploadTransactionModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onSubmit={handleSubmitCsvTransaction}
        isLoading={uploadTransactionMutation.isPending}
      />
      <HeaderPage onClickUpload={() => setIsModalOpen(true)} />
      <Balance balance={balanceQuery.data ?? 0} />
      <TransactionList transactions={unsuccessfulTransactionQuery.data ?? []} />
    </Container>
  );
};

export default Transaction;
