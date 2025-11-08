import { useQuery } from '@tanstack/react-query';
import { RQ_BALANCE } from '@/libs/constant';
import { getBalance } from '@/libs/services/transaction';

function useBalanceQuery() {
  async function getBalanceData() {
    const result = await getBalance();
    if (result.code !== 'SUCCESS') {
      throw new Error(result.message);
    }

    return result.data;
  }

  return useQuery({
    queryKey: [RQ_BALANCE],
    queryFn: getBalanceData,
  });
}

export default useBalanceQuery;
