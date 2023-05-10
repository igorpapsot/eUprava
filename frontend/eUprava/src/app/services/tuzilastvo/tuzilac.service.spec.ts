import { TestBed } from '@angular/core/testing';

import { TuzilacService } from './tuzilac.service';

describe('TuzilacService', () => {
  let service: TuzilacService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TuzilacService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
