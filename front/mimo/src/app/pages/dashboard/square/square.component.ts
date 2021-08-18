import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TurnPageService } from 'src/app/services/turn-page.service';

@Component({
  selector: 'app-square',
  templateUrl: './square.component.html',
  styleUrls: ['./square.component.sass']
})
export class SquareComponent implements OnInit {

  constructor(
    private turnPageService: TurnPageService,
    private route: ActivatedRoute,
    ) { }

  ngOnInit(): void {
  }

    actionToArchive = (aid?: number) => {

        this.turnPageService.goNavigate(['../archive', ], { queryParams: { aid: aid }, relativeTo: this.route})
    }

}
