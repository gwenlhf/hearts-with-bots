import { Component, OnInit } from '@angular/core';
import { UnicardfPipe } from '../unicardf.pipe';
import { RulesService } from '../rules.service';

@Component({
  selector: 'app-flat-board',
  templateUrl: './flat-board.component.html',
  styleUrls: ['./flat-board.component.scss']
})
export class FlatBoardComponent implements OnInit {
	weights : number[];
  constructor(private rulesService: RulesService) { }

  ngOnInit() {
  	this.rulesService.getNewGameFlat()
  		.subscribe(weights => this.weights = weights);
  }

}
