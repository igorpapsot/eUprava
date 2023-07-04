import { TestBed } from '@angular/core/testing';

import { RocisteService } from './rociste.service';

describe('RocisteService', () => {
  let service: RocisteService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RocisteService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
