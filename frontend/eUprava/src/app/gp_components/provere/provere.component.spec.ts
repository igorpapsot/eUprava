import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProvereComponent } from './provere.component';

describe('ProvereComponent', () => {
  let component: ProvereComponent;
  let fixture: ComponentFixture<ProvereComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProvereComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProvereComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
