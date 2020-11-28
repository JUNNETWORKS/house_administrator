import { useState, FC, useEffect } from 'react';
import RoomItem from 'components/roomItem';
import { Room, getAllRooms } from 'domains/administrator';

const AllRooms: FC = () => {
  const [rooms, setRooms] = useState<Room[]>([]);
  useEffect(() => {
    const load = async (): Promise<void> => {
      try {
        const roomDatas = await getAllRooms();
        setRooms(roomDatas);
      } catch (err) {
        throw new Error('Error');
      }
    };

    // 返ってくるPromiseをあえて無視することを明示するためにvoid四季を利用している
    void load();
  }, []);

  return (
    <div>
      {rooms.map((room: Room) => (
        <p>{room.name}</p>
      ))}
    </div>
  );
};

export default AllRooms;
