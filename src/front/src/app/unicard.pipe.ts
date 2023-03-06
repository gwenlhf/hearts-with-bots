import { Pipe, PipeTransform } from '@angular/core';
import { Suit, Value } from './game.constants';
import { Card } from './game.card';
@Pipe({
  name: 'unicard'
})
export class UnicardPipe implements PipeTransform {
  transform(value: Card): string {
  	var uoffset = 126976;
    var valv = (value.Value == Value.Ace) ? 1 : value.Value + 2;
    var suitv;
    switch (value.Suit) {
    	case Suit.Spades:
    		suitv = 10*16;
    		break;
    	case Suit.Hearts:
    		suitv = 11*16;
    		break;
    	case Suit.Diamonds:
    		suitv = 12*16;
    		break;
    	case Suit.Clubs:
    		suitv = 13*16;
    		break;
		default:break;
    }
    var cardv = (uoffset + suitv + valv);
    return String.fromCodePoint(cardv);
  }

}
