import { TestBed } from '@angular/core/testing';

import { SudstvoService } from './sudstvo.service';

describe('SudstvoService', () => {
  let service: SudstvoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SudstvoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
