import React from 'react';

type RoomProps = {
  name: string;
  description: string;
};

const Room: React.FC<RoomProps> = ({ name, description }) => {
  return (
    <>
      <h1>RoomName: {name}</h1>
      <p>{description}</p>
    </>
  );
};

export default Room;
