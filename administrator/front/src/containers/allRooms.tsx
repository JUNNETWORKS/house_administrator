import React from 'react';
import RoomItem from 'components/roomItem';
import { getAllRooms } from 'domains/administrator';

const AllRooms: FC = () => {
  try {
    const allRooms = getAllRooms();
  } catch (e) {
    console.dir(e);
    return <p>Error</p>;
  }
  const roomItems = allRooms.map((room) => <RoomItem>room.

  return (
    <div>
      <RoomItem />;
      </div>
};
