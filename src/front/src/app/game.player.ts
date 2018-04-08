import { Side } from './game.constants';
import { Card } from './game.card';

export class Player {
	Side : Side;
	Hand : Card[];
	Points : Card[];
	Total : number;
}