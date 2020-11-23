import 'tailwindcss/tailwind.css';
import { AppProps } from 'next/app';

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <div className="bg-gray-100">
      <Component {...pageProps} />;
    </div>
  );
};

export default App;
