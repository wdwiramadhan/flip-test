import { useMutation } from '@tanstack/react-query';
import { uploadCsvTransaction } from '@/libs/services/transaction';

function useUploadTransactionMutation() {
  async function uploadTransaction(formData: FormData) {
    const result = await uploadCsvTransaction(formData);
    if (result.code !== 'SUCCESS') {
      throw new Error(result.message);
    }

    return result.data;
  }

  return useMutation({
    mutationFn: uploadTransaction,
  });
}

export default useUploadTransactionMutation;
