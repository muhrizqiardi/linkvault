import { useState } from 'react';

export default function useError(
  initialIsError: boolean,
  initialErrorMessage: string = '',
) {
  const [isError, setIsError] = useState<boolean>(initialIsError);
  const [errorMessage, setErrorMessage] = useState<string>(initialErrorMessage);

  const setError = (isError: boolean, errorMessage: string = '') => {
    setIsError(isError);
    setErrorMessage(errorMessage);
  };

  return {
    isError,
    errorMessage,
    setError,
  };
}
