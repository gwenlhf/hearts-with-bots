import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'heatclass'
})
export class HeatclassPipe implements PipeTransform {

  transform(value: number): string {
  	var val = Number(value.toFixed(1))
    switch (val) {
    	case -1:
    	case -0.9:
    	case -0.8:
    		return "verycold"
    	case -0.7:
    	case -0.6:
    		return "cold"
    	case -0.5:
    	case -0.4:
    	case -0.3:
    		return "cool"
    	case -0.2:
    		return "cold"
    	case -0.1:
    	case 0.0:
    	case 0.1:
    	case 0.2:
    		return "neutral"
    	case 0.3:
    	case 0.4:
    	case 0.5:
    		return "warm"
    	case 0.6:
    	case 0.7:
    		return "hot"
    	case 0.8:
    	case 0.9:
    	case 1.0:
    		return "veryhot"
    	default:
    		return "neutral";
    }
  }
}
