import 'tailwindcss/tailwind.css';
import { AppProps } from 'next/app';
import Header from 'components/header';

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <div className="bg-gray-100">
      <Header title="this is top page" />
      <Component {...pageProps} />;
    </div>
  );
};

export default App;
