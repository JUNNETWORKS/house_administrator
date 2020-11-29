import { useState, FC, useEffect } from 'react';
import { Room, getRoom } from 'domains/administrator';
import RoomDetail from 'components/room';

type RoomProps = {
  id: number;
};

const RoomDetailPage: FC<RoomProps> = ({ id }) => {
  const [room, setRoom] = useState<Room>();
  useEffect(() => {
    const load = async (): Promise<void> => {
      try {
        const roomData = await getRoom(id);
        setRoom(roomData);
      } catch (err) {
        console.dir(err);
        throw new Error('Error');
      }
    };

    // 返ってくるPromiseをあえて無視することを明示するためにvoid式を利用している
    void load();
  }, [id]);

  return (
    <>
      {room ? (
        <RoomDetail name={room.name} description={room.description} />
      ) : (
        <p>NowLoading</p>
      )}
    </>
  );
};

export default RoomDetailPage;
