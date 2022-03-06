#include "fbcgo.hpp"
#include "flatbuffers/flatbuffers.h"
#include "flatbuffers/idl.h"
#include <iostream>

int Create(Parser *p) {
  p->_parser = (void*) new flatbuffers::Parser();
  return 0;
}

int check(Parser *p) {
  if (p == nullptr) {
    return 1;
  }
  if (p->_parser == nullptr) {
    return 1;
  }
  return 0;
}

void Destroy(Parser *p) {
  if (check(p) != 0)
    return;

  delete (flatbuffers::Parser*)(p->_parser);
}

int AddBuffer(Parser *p, char *b) {
  if (check(p) != 0)
    return 1;

  auto pp = (flatbuffers::Parser*)(p->_parser);
  if (!pp->Parse(b)) {
    std::cout << pp->error_ << std::endl;
    return 1;
  }
  return 0;
}

unsigned int GetSize(Parser *p) {
  if (check(p) != 0)
    return 0; /* zero is the only safe answer here */

  auto pp = (flatbuffers::Parser*)(p->_parser);
  return pp->builder_.GetSize();
}

unsigned char *GetBuffer(Parser *p) {
  if (check(p) != 0)
    return nullptr;

  auto pp = (flatbuffers::Parser*)(p->_parser);
  return pp->builder_.GetBufferPointer();
}

void Finish(Parser *p) {
  if (check(p) != 0)
    return;
  auto pp = (flatbuffers::Parser*)(p->_parser);
  pp->builder_.Finish(flatbuffers::Offset<int>(0), nullptr);
}


