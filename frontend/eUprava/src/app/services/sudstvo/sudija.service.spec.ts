import { TestBed } from '@angular/core/testing';

import { SudijaService } from './sudija.service';

describe('SudijaService', () => {
  let service: SudijaService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SudijaService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
