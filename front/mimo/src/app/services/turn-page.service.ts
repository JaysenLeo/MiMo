import { Injectable } from '@angular/core';
import { NavigationExtras, Router } from '@angular/router';

@Injectable({
    providedIn: 'root'
})
export class TurnPageService {

    constructor(
        private router: Router,
    ) {

    }

    goNavigate = (commands: any[], extras?: NavigationExtras) => {
        this.router.navigate(commands, extras);
    }
}
