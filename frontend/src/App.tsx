import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import TransactionModule from '@/modules/Transaction';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <TransactionModule />
    </QueryClientProvider>
  );
};

export default App;
