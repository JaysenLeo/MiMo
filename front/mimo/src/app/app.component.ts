import { Component } from '@angular/core';
import { NzIconService } from 'ng-zorro-antd/icon';
import { Observable } from 'rxjs';
import { NgxAstilectronService } from './share/ngx-astilectron/ngx-astilectron.service';

const ngZorroIconLiteral =
  '<?xml version="1.0" encoding="UTF-8"?><svg width="24" height="24" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M44 35H40L24 6.5L8 35H4" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M24 31C21.2386 31 19 34.5817 19 39V41H29V39C29 34.5817 26.7614 31 24 31Z" fill="none" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M4 41L44 41" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M15 23C15 23 20 19 24 19C28 19 33 23 33 23" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M40 6L38 9L40 12" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M40 6L42 9L40 12" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M7 17L6 19L7 21" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/><path d="M7 17L8 19L7 21" stroke="#333" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/></svg>'


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass']
})
export class AppComponent {
  title = 'mimo';
  RequestValue$!: Observable<any>;
  constructor(
    private iconService: NzIconService,
    private astilectronSrv: NgxAstilectronService
    ) {
    this.iconService.addIconLiteral('ng-zorro:antd', ngZorroIconLiteral);

  }
  ngOnInit(): void {
    //Called after the constructor, initializing input properties, and the first call to ngOnChanges.
    //Add 'implements OnInit' to the class.
    this.astilectronSrv.registerMsgCallback((msg: any) => {
      if(msg.name === 'check.out.menu') {
        console.log('check.out.menu: ', msg.payload);
      }
    })
    this.astilectronSrv.initAstilectron()
    // Wait for astilectron to be ready
    // document.addEventListener('astilectron-ready', () => {
    //   astilectron.onMessage((message: any) => {
    //           switch (message.name) {
    //               case "about":
    //                   console.log(message.payload);
    //                   break;
    //               case "check.out.menu":
    //                   console.log(message.payload);
    //                   break;
    //           }
    //       });
    //   let message = {"name": "hello", "payload":"PayloadTest"};
    //   // asticode.loader.show();
    //   astilectron.sendMessage(message, (message: any) => {
    //     console.log(message)
    //   })
    // })
  }



}
