import { Component, OnInit } from '@angular/core';
import { Game } from '../game';
import { Side } from '../game.constants';
import { GameSidePipe } from '../game-side.pipe';
import { RulesService } from '../rules.service';

@Component({
  selector: 'app-gameboard',
  templateUrl: './gameboard.component.html',
  styleUrls: ['./gameboard.component.scss']
})
export class GameboardComponent implements OnInit {
	Model : Game;
  constructor (private rulesService : RulesService) { }

  ngOnInit() {
  	this.newGame();
  }
  newGame() {
  	this.rulesService.getNewGame()
  		.subscribe(game => this.Model = game);
  }
}
