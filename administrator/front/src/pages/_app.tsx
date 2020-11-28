import 'tailwindcss/tailwind.css';
import Head from 'next/head';
import { AppProps } from 'next/app';
import Header from 'components/header';

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <Head>
        <title>My page</title>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width"
        />
      </Head>
      <div>
        <Header title="this is top page" />
        <Component {...pageProps} />
      </div>
    </>
  );
};

export default App;
