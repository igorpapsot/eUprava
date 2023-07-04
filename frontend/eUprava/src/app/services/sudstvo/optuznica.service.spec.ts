import { TestBed } from '@angular/core/testing';

import { OptuznicaService } from './optuznica.service';

describe('OptuznicaService', () => {
  let service: OptuznicaService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OptuznicaService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
