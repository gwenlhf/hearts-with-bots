import { Component, OnInit } from '@angular/core';

import '../game.constants'
import { GameCardComponent } from '../game-card/game-card.component'
import { GamePlayerComponent } from '../game-player/game-player.component'

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})
export class GameComponent implements OnInit {
	Players : Player[];
	Trick : Card[];
	HeartsBroken : boolean;
	ToPlay : Side;
  constructor() { }

  ngOnInit() {
  }

}
