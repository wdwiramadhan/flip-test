import type { ErrorResponse } from '@/libs/types/apiResponse';

export function serviceErrorHandler(err: unknown): ErrorResponse {
  return {
    code: 'RUNTIME_ERROR',
    message: err instanceof Error ? err.message : 'Oops, something went wrong!',
    data: null,
  };
}
