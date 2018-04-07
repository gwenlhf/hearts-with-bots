import { Component, OnInit } from '@angular/core';
import '../game.constants'
@Component({
  selector: 'app-game-card',
  templateUrl: './game-card.component.html',
  styleUrls: ['./game-card.component.css']
})
export class GameCardComponent implements OnInit {
	Suit : Suit;
	Value : Value;
  constructor() { }

  ngOnInit() {
  }

}
