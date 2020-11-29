import Room from 'containers/room';
import { useRouter } from 'next/router';
import { NextPage } from 'next';

const RoomPage: NextPage = () => {
  const router = useRouter();
  const { id } = router.query;

  if (!id) return <p>query error</p>;

  return <Room id={+id} />;
};
export default RoomPage;
