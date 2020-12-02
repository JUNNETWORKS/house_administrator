import React from 'react';
import { useForm } from 'react-hook-form';
import { createNewRoom } from 'domains/administrator';
import {Room} from '../domains/administrator/models/room';
import { useRouter } from 'next/router';

type formData = {
  name: string;
  description: string;
  HostName: string;
};

const NewRoomForm = () => {
  const { register, handleSubmit, errors } = useForm<formData>();
  const router  = useRouter();
  const onSubmit = async (data: formData) => {
    let newRoomData: Room;
    try{
       newRoomData = await createNewRoom(data);
       router.push(`/rooms/${newRoomData.id}`);
      } catch (err) {
        console.dir(err);
        throw new Error('Error');
      }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <h1>New Room</h1>
      <label>
        Name
        <input name="name" ref={register({ required: true, minLength: 3 })} />
        {errors.name?.type === 'required' && <p>Name is required</p>}
        {errors.name?.type === 'minLength' && <p>minLength is 3</p>}
      </label>
      <label>
        Description
        <textarea name="description" ref={register} />
      </label>
      <label>
        HostName(Domain or IP address)
        <input name="hostName" ref={register({ required: true })} />
        {errors.HostName && <p>HostName is required</p>}
      </label>
      <input type="submit" />
    </form>
  );
};

export default NewRoomForm;
