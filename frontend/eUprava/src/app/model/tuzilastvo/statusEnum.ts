import { Pipe, PipeTransform } from "@angular/core";

export enum Status {
    ODBACENA = 0,
    PRIHVACENA = 1,
    NACEKANJU = 2
}

@Pipe({
    name: 'status',
  })
  export class StatusPipe implements PipeTransform {
    transform(value: any, args?: any): any {
      if (value == Status.ODBACENA) return 'Odbacena';
      if (value == Status.PRIHVACENA) return 'Prihvacena';
      if (value == Status.NACEKANJU) return 'Na cekanju';
      return null;
    }
  }