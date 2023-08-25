import { TestBed } from '@angular/core/testing';

import { GpolicajacService } from './gpolicajac.service';

describe('GpolicajacService', () => {
  let service: GpolicajacService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GpolicajacService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
