import { Component, OnInit, Input } from '@angular/core';

import { Card } from '../game.card';
import { Suit, Value } from '../game.constants';

@Component({
  selector: 'app-game-card',
  templateUrl: './game-card.component.html',
  styleUrls: ['./game-card.component.css']
})
export class GameCardComponent implements OnInit {
	@Input() model : Card;
	Suit : string;
	Value : string;
  ngOnInit() {
  	this.Suit = Suit[this.model.Suit];
  	this.Value = Value[this.model.Value]; 
  }

}
