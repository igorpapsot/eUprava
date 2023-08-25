import { TestBed } from '@angular/core/testing';

import { PrelazakService } from './prelazak.service';

describe('PrelazakService', () => {
  let service: PrelazakService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PrelazakService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
