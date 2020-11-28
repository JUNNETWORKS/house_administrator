import React from 'react';
import Link from 'next/link';

const RoomItem: React.FC = () => {
  return (
    <>
      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          border: '1px solid black',
          width: '20%',
        }}
      >
        <img
          src="https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.thespruce.com%2Fliving-room-lighting-ideas-4134256&psig=AOvVaw2kvz9OKm-sW3NR76ccLkuv&ust=1606579346393000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCLC_082Mo-0CFQAAAAAdAAAAABAD"
          alt="Room Picture"
          width={256}
          height={144}
        />
        <p>Room Name</p>
        <p>Room Description</p>
        <Link href="/rooms/hogehoge">
          <a>Go to Room Page</a>
        </Link>
      </div>
    </>
  );
};

export default RoomItem;
