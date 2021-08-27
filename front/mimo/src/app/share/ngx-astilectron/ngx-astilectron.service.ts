import { ChangeDetectorRef, Injectable } from '@angular/core';
import { BehaviorSubject, Subscription } from 'rxjs';

declare var astilectron: any;
interface msgCallback {
  call: ((msg: any) => any)
}


@Injectable({
  providedIn: 'root'
})
export class NgxAstilectronService {

  callBackPool: msgCallback[]
  messageSendSubject: BehaviorSubject<{name: string, payload: string, cb: (msg: any) => any}>

  constructor(
  ) {
    this.callBackPool = []
    this.messageSendSubject = new BehaviorSubject<{name: string, payload: string, cb: (msg: any) => any}>(
      {
        name: 'hello', payload: 'pong',
        cb: (msg: any) => { console.log('hello recv:', msg) }
      }
    );
   }


  /**
   *
   *
  */
  registerMsgCallback = (messageCallback: (msg: any) => any) => {
    this.callBackPool.push({call: messageCallback});
  }

  initAstilectron = () => {
    // Wait for astilectron to be ready
    document.addEventListener('astilectron-ready', () => {
      astilectron.onMessage((message: any) => {
        console.log('callBackPool:', this.callBackPool)
        this.callBackPool.forEach(callFd => {
            callFd.call(message);
        })
      });
      this.messageSendSubject.subscribe(
        msg => {
          astilectron.sendMessage({name: msg.name, payload: msg.payload}, msg.cb)
        }
      )
    })
  }

  sendMessage = (msg: {name: string, payload: string, cb: (msg: any) => any}) => {
    this.messageSendSubject.next(msg)
  }
}
