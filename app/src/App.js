import './App.css';
import { Plugins } from '@capacitor/core';

const { PushNotifications } = Plugins;

function App() {

  console.log("Init")

  PushNotifications.requestPermission().then( result => {
    if (result.granted) {
      PushNotifications.register()
    } else {
      console.log('Not granted')
    }
  })

  PushNotifications.addListener('registration', token => {
    console.log('Notification granted, token: ', token, ' --- value:', token.value)
  })

  PushNotifications.addListener('registrationError', error => {
    console.log('Error on registration: ', JSON.stringify(error))
  })

  PushNotifications.addListener('pushNotificationReceived', notification => {
    console.log('PUSH RECEIVED: ', JSON.stringify(notification))
  })

  PushNotifications.addListener('pushNotificationActionPerformed', notification => {
    console.log('Push action performed: ', JSON.stringify(notification))
  })


  return (
    <div className='App'>
      <p>Hello world</p>
      <button onClick={() => console.log("hello")}>Click me</button>
    </div>
  );
}

export default App;
