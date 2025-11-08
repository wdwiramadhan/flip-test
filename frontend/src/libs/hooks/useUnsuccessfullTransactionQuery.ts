import { useQuery } from '@tanstack/react-query';
import { RQ_UNSUCCESSFUL_TRANSACTIONS } from '@/libs/constant';
import { getUnsuccessfulTransactions } from '@/libs/services/transaction';

function useUnsuccessfulTransactionQuery() {
  async function getUnsuccessfulTransactionsData() {
    const result = await getUnsuccessfulTransactions();
    if (result.code !== 'SUCCESS') {
      throw new Error(result.message);
    }

    return result.data;
  }

  return useQuery({
    queryKey: [RQ_UNSUCCESSFUL_TRANSACTIONS],
    queryFn: getUnsuccessfulTransactionsData,
  });
}

export default useUnsuccessfulTransactionQuery;
