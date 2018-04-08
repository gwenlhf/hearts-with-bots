import { Side } from './game.constants';
import { Player } from './game.player';
import { Card } from './game.card';

export class Game {
	Players : Player[];
	Trick : Card[];
	HeartsBroken : boolean = false;
	ToPlay : Side;
}