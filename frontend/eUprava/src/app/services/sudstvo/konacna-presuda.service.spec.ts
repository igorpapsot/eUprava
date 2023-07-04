import { TestBed } from '@angular/core/testing';

import { KonacnaPresudaService } from './konacna-presuda.service';

describe('KonacnaPresudaService', () => {
  let service: KonacnaPresudaService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(KonacnaPresudaService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
