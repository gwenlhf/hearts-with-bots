import { Pipe, PipeTransform } from '@angular/core';
import { Suit, Value } from './game.constants'
@Pipe({
  name: 'unicardf'
})
export class UnicardfPipe implements PipeTransform {
  transform(value: number): string {
  	var uoffset = 126976;
    var valv = value % 13;
    valv = (valv == Value.Ace) ? 1 : valv + 2;
    var suitv = Math.floor(value / 13);
    switch (suitv) {
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
