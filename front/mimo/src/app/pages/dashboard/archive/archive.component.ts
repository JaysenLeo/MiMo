import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MarkdownService } from 'ngx-markdown';
import { BehaviorSubject } from 'rxjs/internal/BehaviorSubject';
import { ArchiveRespData } from './archive.interface';
import { ArchiveService } from './archive.service';

@Component({
    selector: 'app-archive',
    templateUrl: './archive.component.html',
    styleUrls: ['./archive.component.css']
})
export class ArchiveComponent implements OnInit {
    Archive$!: BehaviorSubject<ArchiveRespData>;
    Content: string
    constructor(
        private route: ActivatedRoute,
        private archiveService: ArchiveService,
        private markdownService: MarkdownService
    ) {
        this.Content = ''
        this.Archive$ = this.getArchiveObservable()
    }

    ngOnInit(): void {
      console.log(this.route.snapshot.queryParams.aid)
        this.archiveService.getArchive(this.route.snapshot.queryParams.aid)
        this.Archive$.subscribe(archive => {

            this.Content = this.markdownService.compile(archive.content)
            console.log(this.Content)
        })
    }

    getArchiveObservable = (): BehaviorSubject<ArchiveRespData> => this.archiveService.getArchiveObservable()

}
