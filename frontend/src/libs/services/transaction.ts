import { serviceErrorHandler } from '@/libs/helpers/error';
import type { APIResponse } from '@/libs/types/apiResponse';
import type { Transaction } from '@/libs/types/transaction';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export async function getBalance(): Promise<APIResponse<number>> {
  try {
    const response = await fetch(`${API_URL}/transactions/balance`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    return await response.json();
  } catch (error) {
    return serviceErrorHandler(error);
  }
}

export async function getUnsuccessfulTransactions(): Promise<APIResponse<Transaction[]>> {
  try {
    const response = await fetch(`${API_URL}/transactions/issues`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    return await response.json();
  } catch (error) {
    return serviceErrorHandler(error);
  }
}

export async function uploadCsvTransaction(formData: FormData): Promise<APIResponse<null>> {
  try {
    const response = await fetch(`${API_URL}/transactions/upload`, {
      method: 'POST',
      body: formData,
    });

    return await response.json();
  } catch (error) {
    return serviceErrorHandler(error);
  }
}
