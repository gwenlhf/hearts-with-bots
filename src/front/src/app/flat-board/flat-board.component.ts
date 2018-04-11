import { Component, OnInit } from '@angular/core';
import { UnicardfPipe } from '../unicardf.pipe'

@Component({
  selector: 'app-flat-board',
  templateUrl: './flat-board.component.html',
  styleUrls: ['./flat-board.component.scss']
})
export class FlatBoardComponent implements OnInit {
	weights : number[];
  constructor() { }

  ngOnInit() {
  }

}
