import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ErrortodoComponent } from './errortodo.component';

describe('ErrortodoComponent', () => {
  let component: ErrortodoComponent;
  let fixture: ComponentFixture<ErrortodoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ErrortodoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ErrortodoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
