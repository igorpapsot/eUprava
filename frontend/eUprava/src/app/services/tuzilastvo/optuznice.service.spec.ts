import { TestBed } from '@angular/core/testing';

import { OptuzniceService } from './optuznice.service';

describe('OptuzniceService', () => {
  let service: OptuzniceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OptuzniceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
