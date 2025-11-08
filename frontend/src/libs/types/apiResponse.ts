export type SuccessCode = 'SUCCESS';
export type ErrorCode = 'BAD_REQUEST' | 'RUNTIME_ERROR';

export type SuccessResponse<TData> = {
  code: SuccessCode;
  message: string;
  data: TData;
};

export type ErrorResponse = {
  code: ErrorCode;
  message: string;
  data: null;
};

export type APIResponse<TData = null> = SuccessResponse<TData> | ErrorResponse;
