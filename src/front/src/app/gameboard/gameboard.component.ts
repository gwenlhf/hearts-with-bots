import { Component, OnInit } from '@angular/core';
import { Suit, Value } from '../game.constants';
import { Game } from '../game';
import { RulesService } from '../rules.service';

@Component({
  selector: 'app-gameboard',
  templateUrl: './gameboard.component.html',
  styleUrls: ['./gameboard.component.css']
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
