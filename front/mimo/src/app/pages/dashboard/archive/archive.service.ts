import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, EMPTY, Observable, Subscriber, Subscription, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';
import { APIs } from 'src/app/apis/base.api';
import { ArchiveEntity } from './archive';
import { ArchiveRespAPI, ArchiveRespData } from './archive.interface';


@Injectable({
    providedIn: 'root'
})
export class ArchiveService {

    public archiveClient: Subscription
    private ArchiveObj: ArchiveEntity
    constructor(
        private http: HttpClient
    ) {
        this.archiveClient = new Subscriber()
        this.ArchiveObj = new ArchiveEntity()
    }


    getArchiveObj = (): ArchiveRespData => this.ArchiveObj.getArchive()
    getArchiveObservable = (): BehaviorSubject<ArchiveRespData> => this.ArchiveObj.getArchiveObservable()

    getArchive = (aid: number) => {
        this.archiveClient = this.http.get<ArchiveRespAPI>(APIs.archive, {
            headers: {

            },
            params: {
                aid: aid.toString()
            }
        })
            .pipe(
                retry(3), // retry a failed request up to 3 times
                catchError(this.handleError) // then handle the error
            ).subscribe(
                resp => {
                    this.ArchiveObj.setArchive(resp.data)
                },
                err => {},
                () => {
                  this.archiveClient.unsubscribe()
                }
            )
    }

    private handleError(error: HttpErrorResponse) {
        if (error.error instanceof ErrorEvent) {
            // A client-side or network error occurred. Handle it accordingly.
            console.error('An error occurred:', error.error.message);
        } else {
            // The backend returned an unsuccessful response code.
            // The response body may contain clues as to what went wrong.
            console.error(
                `Backend returned code ${error.status}, ` +
                `body was: ${error.error}`);
        }
        // Return an observable with a user-facing error message.
        return throwError(
            'Something bad happened; please try again later.');
    }
}
