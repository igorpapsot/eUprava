import { TestBed } from '@angular/core/testing';

import { PoternicaService } from './poternica.service';

describe('PoternicaService', () => {
  let service: PoternicaService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PoternicaService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
